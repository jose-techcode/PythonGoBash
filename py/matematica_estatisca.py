import math
import statistics
from statistics import mode
import random
import numpy as np

numeros_primos = [2, 3, 3, 5, 7, 11, 13, 17, 19, 23, 29]
numeros_compostos = [4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20]

# Matemática (vetores e matrizes)

print("Soma", (numeros_primos[2] + numeros_primos[5]))
print("Subtração", (numeros_primos[8] - numeros_primos[9]))
print("Multiplicação", (numeros_primos[7] * numeros_primos[4]))
print("Divisão", (numeros_primos[3] / numeros_primos[1]))
print("Potência", (numeros_primos[3] ** numeros_primos[3]))
print("Raiz Quadrada", (numeros_primos[0] ** 0.5))
print("Notação científica", (numeros_primos[0] * 10 ** 10))

# Estatística (média, moda, mediana, máximo, minímmo, variância, desvio padrão, )

print("Média", sum(numeros_primos) / len(numeros_primos))
print("Mediana ímpar", numeros_primos[len(numeros_primos)//2])
print("Mediana par", (numeros_primos[4] + numeros_primos[6]) / 2)
print("Moda", mode(numeros_primos))
print("Máximo", max(numeros_primos))
print("Minímo", min(numeros_primos))
print("Amplitude", max(numeros_primos) - min(numeros_compostos))
print("Variância", statistics.variance(numeros_primos)) # amostral
print("Desvio padrão", statistics.stdev(numeros_primos)) # amostral
print("Coeficiente da variação", statistics.stdev(numeros_primos) / (sum(numeros_primos) / len(numeros_primos)) * 100)

# Distribuição Normal (Estatística)

media = sum(numeros_primos) / len(numeros_primos)

desvio_padrao = statistics.stdev(numeros_primos)

valores_normais = [random.gauss(media, desvio_padrao) for _ in range(10)]

print("Distribuição normal", valores_normais)

# Covariância e Correlação (Estatística)

# Covariância

covariancia = np.cov(numeros_primos, numeros_compostos)[0][1]

# Correlação

correlacao = np.corrcoef(numeros_primos, numeros_compostos)[0][1]

print("Covariância", covariancia)
print("Correlação", correlacao)
