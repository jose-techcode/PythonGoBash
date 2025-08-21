puts ("opa")

puts (10 + 5)

nota = 7

# If & Else

if nota >= 7 then
    puts ("Aprovado")
else
    puts ("Reprovado")
end

# For

for i in 1..3 do
    puts i
end

# While

contador = 0
while contador <= 3 do
    puts contador
    contador += 1
end

# Funções

def saudacao()
    puts "Saudação"
end

saudacao

def soma0(x, y)
    return x + y
end

# puts soma0(5, 5)

# Orientações a Objetos + Funções

class Calculadora
    def soma(a, b)
        return a + b
    end
    
    def subtracao(a, b)
        return a - b
    end

    def multiplicacao(a, b)
        return a * b
    end

    def divisao(a, b)
        return a / b
    end

end

calc = Calculadora.new
puts calc.soma(10, 10)
puts calc.subtracao(10, 10)
puts calc.multiplicacao(10, 10)
puts calc.divisao(10, 10)

class Soma
    def somar(a, b)
        return a + b
    end
end

soma = Soma.new
puts soma.somar(100, 100)

class Subtracao
    def subtrair(a, b)
        return a - b
    end
end

subtracao = Subtracao.new
puts subtracao.subtrair(500, 250)

class Multiplicacao
    def multiplicar(a, b)
        return a * b
    end
end

multiplicacao = Multiplicacao.new
puts multiplicacao.multiplicar(1000, 1000)
        
class Divisao
    def dividir(a, b)
        return a / b
    end
end

divisao = Divisao.new
puts divisao.dividir(2000, 1000)