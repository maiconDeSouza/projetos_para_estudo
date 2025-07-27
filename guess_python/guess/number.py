import random


def random_number(number_init, number_end):
    return random.randint(number_init, number_end)


def compare_numbers_return(hit, msg):
    return {"hit": hit, "msg": msg}


def compare_numbers():
    attempts = 0

    def result(value_project, value_user):
        nonlocal attempts
        if value_project == value_user:
            msg = f"Parabéns você acertou o numero correto é {value_project} em {attempts} tentativas."
            result = compare_numbers_return(True, msg)
            return result

        if value_project > value_user:
            attempts += 1
            msg = f"O Valor é maior que {value_user}."
            result = compare_numbers_return(False, msg)
            return result

        if value_project < value_user:
            attempts += 1
            msg = f"O Valor é menor que {value_user}."
            result = compare_numbers_return(False, msg)
            return result

    return result


def convert_number(input_user):
    return int(input_user)
