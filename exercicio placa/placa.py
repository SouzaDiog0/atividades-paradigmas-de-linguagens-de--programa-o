def inserir_placa():
    placa = input("Digite a placa do seu veículo: ")
    return placa

def validar_placa(placa):
    if len(placa) == 7 and placa[:3].isalpha() and placa[4:].isdigit():
        return True
    else:
        return False
    
def mascarar_placa(placa):
    mascara = placa[:3] + "-" + placa[3:]
    return mascara

def filtro_placa(placa):
    placa =list(placa)
    placa_censurado = placa[:2] + [ '*', '*', '*'] + placa[-2:]
    return ''.join(placa_censurado)

def main():
    placa = inserir_placa()
    placa_valida = validar_placa(placa)
    mascara_placa = mascarar_placa(placa)
    placa_filtro = filtro_placa(placa)
    if placa_valida:
        print("Placa válida!")
        print("Placa do veículo:", mascara_placa)
        print("Placa do veículo com máscara:", placa_filtro)
    else:
        print("Placa inválida!")

if __name__ == "__main__":
    main()