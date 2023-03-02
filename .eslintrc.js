module.exports = {
  root: true,
  env: {
    browser: true,
    es2021: true
  },
  extends: ['airbnb', 'airbnb-typescript', 'airbnb/hooks', 'plugin:react/recommended', 'prettier'],
  parserOptions: {
    ecmaVersion: 'latest',
    sourceType: 'module',
    project: ['web/tsconfig.json', 'docs/tsconfig.json'],
    ecmaFeatures: {
      jsx: true
    },
    tsconfigRootDir: __dirname
  },
  plugins: ['react', 'prettier', 'simple-import-sort', 'import'],
  rules: {
    'react/react-in-jsx-scope': 'off',
    'import/prefer-default-export': 'off',
    'react/jsx-props-no-spreading': 'off', // switched off for component building
    'react/require-default-props': 'off',
    'react/jsx-filename-extension': [
      1,
      {
        extensions: ['.tsx', '.ts']
      }
    ],
    'react/function-component-definition': [
      1,
      {
        namedComponents: 'arrow-function'
      }
    ],
    'react/no-unknown-property': [
      'error',
      {
        ignore: ['jsx']
      }
    ],
    'simple-import-sort/exports': 'warn',
    'simple-import-sort/imports': [
      'warn',
      {
        groups: [
          ['^node:'],
          // Packages `react` related packages
          ['^react', '^next', '^@?\\w'],
          ['@patternfly'],
          ['^@app'],
          // Internal packages.
          ['^~(/.*|$)'],
          // Relative imports
          ['^\\.\\.(?!/?$)', '^\\.\\./?$', '^\\./(?=.*/)(?!/?$)', '^\\.(?!/?$)', '^\\./?$'],
          // Style imports.
          ['^.+\\.?(css|scss)$']
        ]
      }
    ]
  },
  settings: {
    'import/resolver': {
      typescript: {
        project: ['./tsconfig.json']
      }
    }
  }
};
