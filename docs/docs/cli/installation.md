# Installation

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

<Tabs>
<TabItem value="mac" label="MacOS" default>
Use <a>brew</a> package manager
<br/>
<br/>

Add the 1-platform brew tap

```bash
brew tap 1-platform/homebrew-tools
```

Now you can install the CLI

```bash
brew install 1-platform/apic
```

</TabItem>
<TabItem value="alpine" label="Alpine" >
We use cloudsmith to distribute packages
<br/>
<br/>

Add 1-platform cloudsmith repo url to your package manager

```bash
apk add --no-cache bash curl

curl -1sLf \
  'https://dl.cloudsmith.io/public/1-platform/apic/cfg/setup/bash.alpine.sh' \
  | bash
```

Now you can install the CLI

```bash
apk add apic --update-cache
```

</TabItem>

<TabItem value="deb" label="Debian/Ubuntu" >
We use cloudsmith to distribute packages
<br/>
<br/>

Add 1-platform cloudsmith repo url to your package manager

```bash
apt-get update && apt-get install -y bash curl

curl -1sLf \
  'https://dl.cloudsmith.io/public/1-platform/apic/cfg/setup/bash.deb.sh' \
  | bash
```

Now you can install the CLI

```bash
apt-get update && apt-get install -y apic
```

</TabItem>

<TabItem value="linux" label="RedHat/Amazon-Linux" >
We use cloudsmith to distribute packages
<br/>
<br/>

Add 1-platform cloudsmith repo url to your package manager

```bash
curl -1sLf \
  'https://dl.cloudsmith.io/public/1-platform/apic/cfg/setup/bash.rpm.sh' \
  | sh
```

Now you can install the CLI

```bash
yum install apic
```

If your a fedora user

```bash
dnf install apic
```

</TabItem>

<TabItem value="widows" label="Windows" >
We use scoop to distribute packages
<br/>
<br/>

Add 1-platform scoop repo url

```bash
scoop bucket add https://github.com/1-Platform/scoop-tools
```

Now you can install the CLI

```bash
scoop install apic
```

</TabItem>

</Tabs>

To make sure it's installed correctly you can run

```bash
apic --version
```

### Run against an OpenAPI

Let's run against [petstore api](https://google.com)

```bash
apic run --schema https://petstore3.swagger.io/api/v3/openapi.json
```
