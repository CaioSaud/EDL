n=2
primeiro = int(input("Digite o 1º número do intervalo que você deseja: "))
final =  int(input("Digite o 1º número do intervalo que você deseja: "))
if primeiro > final:
    print("Primeiro número tem que ser maior que o final")
else:
    for i in range (primeiro,final+1):
        for t in range (2,i//2+1):
            if i%2==0:
                break
            elif i%t==0:
                break
            if t==i//2:
                n=n+1
print (n)
