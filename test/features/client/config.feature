Feature: Config
  Config allows the system to download a new configuration.

  Scenario Outline: Download existing config
    When I download the "<name>" configuration
    Then I should have a configuration

    Examples: With different configs
      | name |
      | grpc |
      | http |

  Scenario Outline: Download missing config
    When I download the "<name>" configuration
    Then I should not have a configuration
    And I should see a log entry of "version not found" in the file "reports/config.log"

    Examples: With different configs
      | name         |
      | invalid_grpc |
      | invalid_http |
