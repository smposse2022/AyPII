# Guía Heap (Montículos)
## Ejercicios
### Parte I: Seguimientos

Dado el siguiente montículo de MÁXIMO (del cual sólo se muestran las prioridades de sus elementos):
```
         [25]
    [15]      [20]
  [6]  [10]
```

1. Ejecutar la siguiente secuencia: Insertar 45, Insertar 22, Remover, Insertar 5. Realice los dibujos del montículo luego de cada momento (tanto el árbol como el arreglo).

   Puede verificar sus pasos en visualgo.net/en/heap

2. En un montículo de MÍNIMO vacío, ejecutar la secuencia (de la cual sólo conocemos las prioridades de los elementos):

   10, 12, 1, 14, 6, 5, 8, 15, 3, 9, 7, 4, 11, 13, 2. Remover, Remover.


### Parte II: Modificando Montículos

Dado un Montículo de MÁXIMO:

3. Escriba un constructor alternativo `NuevoMonticuloMaxDesdeArreglo` que tome como parámetro un arreglo desordenado de elementos y devuelva un montículo que contenga esos elementos. Escriba los tests correspondientes (al menos 3).

4. Escriba una función que, dado un Montículo de tamaño M y un número entero n, con n >= 1 y n <= M,
devuelva el enésimo máximo del montículo (Ej: si n == 1, devolver el primer máximo). Escriba los tests correspondientes (al menos 3).

5. Escriba una función que, dados dos montículos, devuelva un tercero que sea la combinación de ambos. Escriba los tests correspondientes (al menos 3).

### Ejercicios extra

6. Resuelva los puntos anteriores de la Parte II para la Montículo de MÍNIMO.

7. **Implementar una cola de prioridad en una clínica de emergencias**. En la bulliciosa clínica de emergencias de la ciudad, cada segundo cuenta y la eficiencia es crucial para salvar vidas. La cola de prioridad permitirá gestionar de manera óptima la atención de los pacientes, quienes llegan con distintas urgencias. Cada paciente será asignado un número entero que representará su nivel de prioridad, donde un valor mayor indicará una urgencia más crítica. La cola debe tener la capacidad de agregar nuevos pacientes y garantizar que aquellos con la prioridad más alta sean atendidos primero. Para implementar esta cola de prioridad, se utilizará un montículo genérico, asegurando así un rendimiento eficiente. Además, se deben escribir al menos tres pruebas detalladas para verificar que el sistema funcione correctamente y pueda manejar la afluencia constante de pacientes en situación de emergencia.
