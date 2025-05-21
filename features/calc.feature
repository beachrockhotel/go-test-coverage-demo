Feature: Adding numbers

  Scenario: Add two numbers
    Given the first number is 10
    And the second number is 5
    When I add the numbers
    Then the result should be 15
