# Sygna Bridge IVMS Utils
The IVMS util is an open source library to help you generate the data for interVASP messaging within Sygna Bridge Ecosystem.

> Please navigate to language specific packages to watch detailed description.

## If you want to generate this util for you language
* Navigate to https://github.com/OpenAPITools/openapi-generator and install openapi code generator in your system
* Execute openapi-codegen using `bridge_and_ivms.json` as input
* For example:
  ```
  java -jar openapi-generator-cli.jar \
    generate -i path/to/bridge_and_ivms.json \
    -g typescript-fetch \
    -t /path/to/custom/template \
    -o /path/to/output
  ```
* If generated code does not fit your requirements, you can get  language-specific template and edit yourself
  ```
  java -jar openapi-generator-cli.jar \
    author template -g typescript-fetch \
    -o /path/to/template/output
  ```
* For more information about code generator, visit https://openapi-generator.tech/

> IVMS definition is inspired by https://github.com/steegi/trisa/blob/ivms101_git_34/proto/ivms101/ivms101.proto