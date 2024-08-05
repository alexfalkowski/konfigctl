Feature: Config

  Config allows the system to download a new configuration.

  Scenario: Download existing config
    When I download the configuration
    Then I should have a configuration
