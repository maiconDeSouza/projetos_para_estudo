from guess import number, message


def go():
    number_game = number.random_number(1, 100)
    compare_numbers = number.compare_numbers()
    print("Olá tente acertar o numéro secreto!")

    while True:
        print("")
        input_user = input("Digite seu numero: ")
        try:
            number_user = number.convert_number(input_user)
        except Exception:
            message.show(
                f"Digite apenas números inteiros e valídos. {input_user} não é número inteiro"
            )
            continue
        compare_numbers_result = compare_numbers(number_game, number_user)
        if compare_numbers_result["hit"]:
            message.show(compare_numbers_result["msg"])
            break
        message.show(compare_numbers_result["msg"])
