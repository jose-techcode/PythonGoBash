puts "Digite seu nome:"
nome = gets.chomp   # lê a linha digitada e remove a quebra de linha, além de naturalmente ser string

puts "Olá, #{nome}."

puts "Digite sua idade:"
idade = gets.chomp.to_i  # lê e converte para inteiro

puts "Você tem #{idade} anos."

puts "Digite a nota do aluno:"
nota = get.chomp.to_f # lê e converte para decimal

puts "A nota do aluno é #{nota}."