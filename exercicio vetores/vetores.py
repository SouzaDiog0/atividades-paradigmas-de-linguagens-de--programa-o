def main():
  vetor_a =[]
  vetor_b =[]
  vetor_c =[]
  for i in range(10):
    vetor_a.append(int(input("Digite o elemento %d do vetor A: " %i)))
    vetor_b.append(int(input("Digite o elemento %d do vetor b: " %i)))

  # The second for loop was outside the main function. It should be inside.
  for i in range(10):
      vetor_c.append(vetor_a[i] + vetor_b[i])

  print("Vetor C:")
  for i in range(10):
    print((i, vetor_c[i]))

if __name__ == "__main__":
  main()
