package teams

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type TeamStore struct {
	mongoDriver                  *mongo.Database
	team                         *mongo.Collection
	teamMembership               *mongo.Collection
	teamCollectionName           string
	teamMembershipCollectionName string
}

func NewStore(db *mongo.Database) Store {
	var teamMembershipCollectionName = "team-memberships"
	var teamCollectionName = "teams"

	_, err := db.ListCollectionNames(context.Background(), bson.M{"name": teamCollectionName})
	if err != nil {
		colation := &options.Collation{Locale: "en"}
		opts := options.CreateCollection().SetCollation(colation)
		if err := db.CreateCollection(context.Background(), teamCollectionName, opts); err != nil {
			panic(err)
		}
	}

	return &TeamStore{
		mongoDriver:                  db,
		teamMembership:               db.Collection(teamMembershipCollectionName),
		team:                         db.Collection(teamCollectionName),
		teamCollectionName:           teamCollectionName,
		teamMembershipCollectionName: teamMembershipCollectionName,
	}
}

func (ts *TeamStore) GetATeamById(ctx context.Context, id primitive.ObjectID) (*Team, error) {
	var doc Team
	filter := bson.M{"_id": id}
	err := ts.team.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}

func (ts *TeamStore) GetATeamBySlug(ctx context.Context, slug string) (*Team, error) {
	var doc Team
	filter := bson.M{"slug": slug}
	err := ts.team.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}

func (ts *TeamStore) InsertTeam(ctx context.Context, team *Team) (id primitive.ObjectID, err error) {
	res, err := ts.team.InsertOne(ctx, team)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (ts *TeamStore) UpdateTeamById(ctx context.Context, userId primitive.ObjectID, teamId primitive.ObjectID, team *UpdateTeamDTO) (*Team, error) {
	doc := make(bson.M)
	if team.Name != nil {
		doc["name"] = team.Name
	}
	if team.Description != nil {
		doc["description"] = team.Description
	}
	if team.Slug != nil {
		doc["slug"] = team.Slug
	}
	doc["updated_by"] = userId
	doc["updated_at"] = time.Now()

	var updatedTeam Team
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	if err := ts.team.FindOneAndUpdate(ctx, bson.M{"_id": teamId}, bson.M{"$set": doc}, opts).Decode(&updatedTeam); err != nil {
		return nil, err
	}

	return &updatedTeam, nil
}

func (ts *TeamStore) DeleteTeamById(ctx context.Context, teamId primitive.ObjectID) (*Team, error) {
	var deletedTeam Team
	if err := ts.team.FindOneAndDelete(ctx, bson.M{"_id": teamId}).Decode(&deletedTeam); err != nil {
		return nil, err
	}

	_, err := ts.teamMembership.DeleteMany(ctx, bson.M{"team_id": teamId})
	if err != nil {
		return nil, err
	}

	return &deletedTeam, nil
}

func (ts *TeamStore) ListTeams(ctx context.Context, dto *ListTeamDTO) (*Pagination[[]Team], error) {
	var res []Team
	// various options
	opts := options.Find().SetSkip(dto.Offset).SetLimit(dto.Limit)
	if dto.SortDir == "asc" {
		opts.SetSort(bson.D{bson.E{Key: dto.Sort, Value: 1}})
	} else {
		opts.SetSort(bson.D{bson.E{Key: dto.Sort, Value: -1}})
	}

	cursor, err := ts.team.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &res); err != nil {
		return nil, err
	}
	// estimated count total
	count, err := ts.team.EstimatedDocumentCount(ctx)
	if err != nil {
		panic(err)
	}

	return &Pagination[[]Team]{Total: count, Records: res}, nil
}

func (ts *TeamStore) ListTeamMembers(ctx context.Context, teamId primitive.ObjectID, dto *ListTeamMembersDTO) (*Pagination[[]TeamMemberDetailed], error) {
	// as we need to lookup this will be a aggregation pipeline
	skipStage := bson.D{{Key: "$skip", Value: dto.Offset}}
	var sortStage bson.D
	if dto.SortDir == "asc" {
		sortStage = append(sortStage, bson.E{Key: dto.Sort, Value: 1})
	} else {
		sortStage = append(sortStage, bson.E{Key: dto.Sort, Value: -1})
	}
	limitStage := bson.D{{Key: "$limit", Value: dto.Limit}}
	populateStage := bson.D{{
		Key:   "$lookup",
		Value: bson.M{"from": "auth", "localField": "user_id", "foreignField": "_id", "as": "user"},
	}}
	unwindStage := bson.D{{Key: "$unwind", Value: "$user"}}
	filterStage := bson.D{{Key: "$match", Value: bson.D{{Key: "team_id", Value: teamId}}}}

	facetStage := bson.D{{Key: "$facet", Value: bson.D{
		{Key: "records", Value: bson.A{skipStage, limitStage, populateStage, unwindStage}},
		{Key: "total", Value: bson.A{bson.D{{Key: "$count", Value: "count"}}}},
	}}}

	cursor, err := ts.teamMembership.Aggregate(ctx, mongo.Pipeline{filterStage, facetStage})
	if err != nil {
		return nil, err
	}
	var doc AggregatedPagination[[]TeamMemberDetailed]
	if cursor.Next(ctx) {
		if err = cursor.Decode(&doc); err != nil {
			return nil, err
		}
	}

	return &Pagination[[]TeamMemberDetailed]{
		Total:   doc.Total[0].CountDocuments,
		Records: doc.Records,
	}, nil
}

// team membership operations
func (ts *TeamStore) InsertTeamMembership(ctx context.Context,
	membership *TeamMember) (id primitive.ObjectID, err error) {
	res, err := ts.teamMembership.InsertOne(ctx, membership)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (ts *TeamStore) RemoveTeamMembership(ctx context.Context, teamId primitive.ObjectID, membershipId primitive.ObjectID) error {
	_, err := ts.teamMembership.DeleteOne(ctx, bson.M{"_id": membershipId, "team_id": teamId})
	if err != nil {
		return err
	}

	return nil
}

func (ts *TeamStore) GetTeamMemberCount(ctx context.Context, teamId primitive.ObjectID) (int64, error) {
	filter := bson.M{"team_id": teamId}
	count, err := ts.teamMembership.CountDocuments(ctx, filter)
	if err != nil {
		return 0, nil
	}

	return count, nil
}

func (ts *TeamStore) GetTeamMembershipById(ctx context.Context, membershipId primitive.ObjectID) (membership *TeamMember, err error) {
	var res TeamMember
	err = ts.teamMembership.FindOne(ctx, bson.M{"_id": membershipId}).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (ts *TeamStore) GetTeamMembershipByEmail(ctx context.Context, teamId primitive.ObjectID, email string) (membership *TeamMemberDetailed, err error) {
	filterStage := bson.D{{Key: "$match", Value: bson.D{{Key: "team_id", Value: teamId}}}}
	populateUser := bson.D{{
		Key:   "$lookup",
		Value: bson.M{"from": "auth", "localField": "user_id", "foreignField": "_id", "as": "user"},
	}}
	unwindStage := bson.D{{Key: "$unwind", Value: bson.D{
		{Key: "path", Value: "$user"},
		{Key: "preserveNullAndEmptyArrays", Value: true},
	}}}
	filterByEmailStage := bson.D{{Key: "$match", Value: bson.D{{Key: "$or", Value: bson.A{
		bson.D{{Key: "user.email", Value: email}},
		bson.D{{Key: "invitee_email", Value: email}},
	}}}}}
	cursor, err := ts.teamMembership.Aggregate(ctx, mongo.Pipeline{filterStage, populateUser, unwindStage, filterByEmailStage})
	if err != nil {
		return nil, err
	}

	var doc []TeamMemberDetailed
	if err = cursor.All(ctx, &doc); err != nil {
		return nil, err
	}

	if len(doc) < 1 {
		return nil, mongo.ErrNoDocuments
	}

	return &doc[0], nil
}

func (ts *TeamStore) UpdateTeamMembershipById(ctx context.Context, membershipId primitive.ObjectID,
	tm *UpdateTeamMembershipDTO) (err error) {
	doc := bson.M{}
	if tm.Email != nil && *tm.Email == "" {
		// remove the invitee field from the collection when time to remove as its onetime field
		doc["$unset"] = bson.M{"invitee_email": 1}
	}
	setDoc := make(bson.M)
	if tm.UserId != nil {
		setDoc["user_id"] = tm.UserId
	}
	if tm.Role != nil {
		setDoc["role"] = tm.Role
	}
	if tm.Status != nil {
		setDoc["status"] = tm.Status
	}
	doc["$set"] = setDoc
	// add more as required by each operation
	res, err := ts.teamMembership.UpdateOne(ctx, bson.M{"_id": membershipId}, doc)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (ts *TeamStore) GetTeamMembershipByUserId(ctx context.Context, teamId primitive.ObjectID, userId primitive.ObjectID) (membership *TeamMember, err error) {
	var res TeamMember
	err = ts.teamMembership.FindOne(ctx, bson.M{"team_id": teamId, "user_id": userId}).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
