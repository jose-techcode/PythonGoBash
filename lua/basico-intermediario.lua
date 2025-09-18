print("Olá mundo")

print(10 - 5)

-- Lógica booleana

local x = 5
local y = 10

-- AND

if x == 5 and y == 10 then
    print("X é igual a 5 e Y é igual 10.")
else
    print("Tudo falso.")
end

-- OR

if x == 10 or y == 10 then
    print("X ou Y é igual a 10.")
else
    print("Nada é igual a 10.")
end

-- NOT AND

if not x == 10 and not y == 5 then
    print("Tudo verdadeiro.")
else
    print("Tudo falso.")
end

-- NOT OR

if not x == 5 or not y == 10 then
    print("X não é igual a 5 e Y não é igual a 10.")
else
    print("X é igual a 5 e Y é igual a 10.")
end

-- Estrutura de controle (then e done)

-- INTEIRO/DECIMAL

local nota = 6
if nota >= 7 then
    print("Passou")
else
    print("Não passou")
end

-- STRING

local login = "jose"
if login == "jose" then
    print("Acessado")
else
    print("Negado")
end

-- FOR (UM EM UM)

for i = 1, 4 do
    print(i)
end

-- FOR (DOIS EM DOIS)

for i = 1, 4, 2 do
    print(i)
end

-- Estrutura de dados

-- LISTA

local numeros = {1, 2, 3}

for i, valor in ipairs(numeros) do -- ipairs para vetores
    print(valor)
end

print(numeros[1])

-- DICIONÁRIO

local pessoa = {
    nome = "José",
    idade = 25,
    cidade = "São Paulo"
}

for conteudo, dados in pairs(pessoa) do -- pairs para string e numero
    print(conteudo, dados)
end

print(pessoa.nome)
print(pessoa["cidade"])

-- Funções

function Saudacao(nome)
    print("Olá," ..nome)
end

Saudacao("jose")

function Idade(numero)
    print("Idade:" ..numero)
end

Idade(27)

function Pessoa(nome, idade)
    print("Nome:" .. nome .. "|" .. "Idade:" .. idade)
end

Pessoa("Jose", 37)

function Soma(x, y)
    print("Resultado:" ..x + y)
end

Soma(5, 5)

function Subtracao(x, y)
    print("Resultado:" ..x - y)
end

Subtracao(10, 5)

function Multiplicacao(x, y)
    print("Resultado:" ..x * y)
end

Multiplicacao(10, 10)

function Divisao(x, y)
    print("Resultado:" .. x / y)
end

Divisao(100, 10)