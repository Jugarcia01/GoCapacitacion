  
Feature: Retirar un monto de la cuenta
  Scenario:
    Given la cuenta tiene un monto de 10000
    When se retira 2000
    Then el monto de la cuenta es de 8000

  Scenario:
    Given la cuenta tiene un monto de 3000
    When se retira 5000
    Then el monto de la cuenta es de 3000
    And genera error "El valor excede el monto actual de la cuenta"

  Scenario:
    Given la cuenta tiene un monto de 3000
    When se retira -1000
    Then el monto de la cuenta es de 3000
    And genera error "El valor no debe ser negativo"