Feature: Buscar las mascotas con sus datos

  Scenario: Buscar mascotas con 4 registros
    Given Un listado de mascotas con 4 registros
    When se solicita buscar todas las mascotas
    Then aparece el listado de 4 mascotas

  Scenario: Buscar mascota con nombre Tony
    Given Buscar mascota con el nombre Tony
    When se ingresa el nombre Tony
    Then aparecen registros de mascotas con el nombre Tony

  Scenario: Buscar mascota por nombre NeverLookIt 
    Given Buscar mascota por el nombre de mascota
    When se busca con la palabra NeverLookIt
    Then aparecen 0 registros de mascotas
    And genera error "No existe ningun registro de mascota con la palabra ingresada"


  Scenario: Buscar mascota por tipo palabra perro
    Given Buscar mascota por el tipo de mascota
    When se busca con la palabra perro
    Then aparecen los registros de mascotas de tipo perro

  Scenario: Buscar mascota por tipo carnivorus
    Given Buscar mascota por el tipo de mascota
    When se busca con la palabra carnivorus
    Then aparecen 0 registros de mascotas
    And genera error "No existe ningun registro de mascota con la palabra ingresada"

  Scenario: Buscar mascota por Id num y letras
    Given Buscar mascota por el ID de mascota
    When se busca con un codigo de letras y numeros de 4 caracteres
    Then aparece registro de mascota con el ID ingresado

  Scenario: Buscar mascota por Id que no existe
    Given Buscar mascota por un ID de mascota que no existe en los registros
    When se busca con un codigo de letras y numeros de 4 caracteres
    Then No aparece ningun registro 
    And genera error "No existe ningun registro de mascota con la palabra ingresada"

