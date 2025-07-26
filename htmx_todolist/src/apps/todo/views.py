from django.shortcuts import render

from .models import Todo


# Create your views here.
def home(request):
    todos = Todo.objects.all().order_by('created_at')
    return render(request, 'todo/pages/home.html', {'todos': todos})
