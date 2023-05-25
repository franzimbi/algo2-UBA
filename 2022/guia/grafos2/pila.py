class Pila:
    """ Crea una pila vacía. """

    def __init__(self):
        self.items = []

    """ Agrega el elemento x a la pila. """

    def apilar(self, x):
        self.items.append(x)

    """ Devuelve el elemento tope y lo elimina de la pila.
        Si la pila está vacía levanta una excepción. """

    def desapilar(self):
        try:
            return self.items.pop()
        except IndexError:
            raise ValueError("La pila está vacía")

    """ Devuelve True si la lista está vacía, False si no. """

    def es_vacia(self):
        return self.items == []
