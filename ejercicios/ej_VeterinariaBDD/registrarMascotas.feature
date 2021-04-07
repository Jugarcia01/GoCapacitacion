Feature: Registrar mascota con los datos necesarios de ID, nombre, tipoMascota, edad (en meses), peso (en libras)

  Scenario: Registrar una mascota
    Given Lista de mascotas es vacia
    When Cree una mascota con los datos necesarios
    Then Lista tiene una mascota

  Scenario: Registrar una mascota con datos correctos
    Given Registar mascota con nombre Danger, tipo de mascota Perro, edad 10 meses y peso 45 libras
    When Se indique crear mascota
    Then Lista de registros tiene mascota con un Id dado, nombre Danger, tipo de mascota Perro, edad 10 meses y peso 45 libras


  Scenario: Intentar registar una mascota sin nombre
    Given Lista de mascotas tiene 3 registros 
    When Registar mascota con nombre vacio
    Then Lista de registros de mascotas tiene 3 registros
    And  genera error "Digite el nombre de la mascota"

  Scenario: Intentar registrar mascota con numeros en el nombre 
    Given Lista de mascotas tiene 3 registros 
    When Registar mascota con nombre 654321
    Then Lista de registros de mascotas tiene 3 registros
    And  genera error "Colocar solamente letras en el nombre de la mascota"


  Scenario: Intentar registrar mascota sin edad
    Given Lista de mascotas tiene 3 registros 
    When Registar mascota sin ingresar edad
    Then Lista de registros tiene mascota 3 registros
    And  genera error "Ingrese la edad de la mascota"

  Scenario: Intentar registrar una mascota con edad en letras
    Given Lista de mascotas tiene 3 registros 
    When Registar mascota con edad en meses: DIEZ
    Then Lista de registros tiene mascota 3 registros
    And  genera error "Colocar solo numeros en la edad de la mascota"

  Scenario: Intentar registrar una mascota con valor negativo de edad 
    Given Lista de mascotas tiene 3 registros 
    When Registar mascota con edad en meses: -10
    Then Lista de registros tiene mascota 3 registros
    And  genera error "La edad de la mascota no puede ser menor o igual a 0"

  Scenario: Intentar registrar una mascota con valor de edad 300
    Given Lista de mascotas tiene 3 registros 
    When Registar mascota con edad en meses: 300
    Then Lista de registros tiene mascota 3 registros
    And  genera error "La edad máxima es 220 meses"


  Scenario: Intentar registrar mascota sin peso
    Given Lista de mascotas tiene 3 registros 
    When Registar mascota sin ingresar peso
    Then Lista de registros tiene mascota 3 registros
    And  genera error "Ingrese el peso de la mascota"

  Scenario: Intentar registrar una mascota con peso en letras
    Given Lista de mascotas tiene 3 registros 
    When Registar mascota con peso en kg: CUARENTA
    Then Lista de registros tiene mascota 3 registros
    And  genera error "Colocar solo numeros en el peso de la mascota"

  Scenario: Intentar registrar mascota con peso negativo
    Given Lista de mascotas tiene 3 registros 
    When Registar mascota con peso en kg: -10
    Then Lista de registros tiene mascota 3 registros
    And  genera error "El peso de la mascota no puede ser menor o igual a 0"

  Scenario: Intentar registrar mascota con peso en libras mayor a 80
    Given Lista de mascotas tiene 3 registros 
    When Registar mascota con peso en libras: 100
    Then Lista de registros tiene mascota 3 registros
    And  genera error "El peso máximo es 80 libras"
