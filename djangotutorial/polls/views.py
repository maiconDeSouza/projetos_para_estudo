from django.shortcuts import render, redirect, get_object_or_404
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


class Details(View):
    def get(self, request, pk):
        question = get_object_or_404(Question, pk=pk)

        context = {
            'question': question,
        }

        return render(request, 'polls/pages/details.html', context)
