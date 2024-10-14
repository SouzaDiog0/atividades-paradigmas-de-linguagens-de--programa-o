class Configuracao:
    _instancia = None

    def __new__(cls, *args, **kwargs):
        if cls._instancia is None:
            cls._instancia = super().__new__(cls, *args, **kwargs)
        return cls._instancia

# Exemplo de uso
config1 = Configuracao()
config2 = Configuracao()

print(config1 is config2)  # Saída: True (mesma instância)
