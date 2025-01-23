Feature: Config
  Config allows the system to download a new configuration.

  Scenario: Download existing config
    When I download the configuration
    Then I should have a configuration

  Scenario: Download missing config
    When I download a missing configuration
    Then I should not have a configuration
    And I should see a log entry of "version not found" in the file "reports/config.log"
