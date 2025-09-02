import os
import math


class Word:
    def __init__(self, word, tip):
        self.word = word.lower()
        self.tip = tip

    def length(self):
        return len(self.word)

    def contains(self, letter):
        return letter in self.word

    def render(self, hits):
        return ' '.join(l if l in hits else '_' for l in self.word)

    def is_solved(self, hits):
        return all(l in hits for l in self.word)


class PlayerGuesses:
    def __init__(self):
        self.hits = set()
        self.errors = set()

    def register(self, letter, word):
        if word.contains(letter):
            self.hits.add(letter)
        else:
            self.errors.add(letter)

    def attempts(self):
        return len(self.hits) + len(self.errors)


class Game:
    def __init__(self, word):
        self.word = word
        self.guesses = PlayerGuesses()
        self.max_attempts = math.ceil(self.word.length() * 1.8)

    def play_turn(self, letter):
        self.guesses.register(letter, self.word)

    def is_over(self):
        return (
            self.word.is_solved(self.guesses.hits)
            or self.guesses.attempts() >= self.max_attempts
        )

    def is_won(self):
        return self.word.is_solved(self.guesses.hits)


class Screen:
    @staticmethod
    def clear():
        os.system('cls' if os.name == 'nt' else 'clear')

    @staticmethod
    def show(game):
        print(game.word.tip)
        print(f'\nCom {game.word.length()} letras')
        print(f'\nErros: {", ".join(game.guesses.errors)}')
        print(game.word.render(game.guesses.hits))
