from django.shortcuts import render, redirect
from django.views import View

from .models import Question


# Create your views here.
class Initial(View):
    def get(self, request):
        return redirect('polls:index')


class Index(View):
    def get(self, request):
        questions = Question.objects.all()

        context = {
            'questions': questions,
        }

        return render(request, 'polls/pages/index.html', context)
