class StripePayment:
    def pay(self, amount):
        print(f"pagamento de ${amount} via stripe processado.")


class PaypalPayment:
    def send_payment(self, amount):
        print(f"pagamento de ${amount} via Paypal processado.")


 #Adapter para adaptar a interface do Paypal ao Stripe

class PaypalAdapter(StripePayment): # Fixed indentation
    def __init__ (self, paypal_payment):
        self.paypal_payment = paypal_payment

    def pay(self, amount): #Fixed indentation
        self.paypal_payment.send_payment(amount)

#Uso
def process_payment(payment_system, amount):
    payment_system.pay(amount) # Added missing function call to process payment

#Exemplo de uso

stripe_payment = StripePayment() # Changed variable name to avoid overwriting the class name
paypal_payment = PaypalAdapter(PaypalPayment())

process_payment(stripe_payment, 100) #pagamento via stripe
process_payment(paypal_payment, 200) #pagamento via paypal