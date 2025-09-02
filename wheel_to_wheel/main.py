import random

from project.run import Game, Word, Screen
from word_list import word_list


def main():
    chosen = random.choice(word_list)
    game = Game(Word(chosen['word'], chosen['tip']))

    while not game.is_over():
        Screen.show(game)
        letter = input('\nDigite uma letra: ').lower().strip()
        game.play_turn(letter)

    Screen.clear()
    if game.is_won():
        print(
            f'🎉 Parabéns, você venceu em {game.guesses.attempts()} tentativas!'
        )
    else:
        print(f"❌ Você perdeu! A palavra era '{game.word.word}'.")


if __name__ == '__main__':
    main()
