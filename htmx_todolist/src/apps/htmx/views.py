from urllib.parse import parse_qs

from django.shortcuts import render, HttpResponse, redirect
from django.shortcuts import get_object_or_404
from django.views.decorators.csrf import csrf_exempt
from django.utils.decorators import method_decorator
from django.views import View

from apps.todo.models import Todo


@method_decorator(csrf_exempt, name='dispatch')
class CRUD(View):
    def post(self, request):
        task = request.POST.get('task')
        if task:
            todo = Todo.objects.create(task=task)
            return render(
                request, 'todo/partials/list-task.html', {'todo': todo}
            )

    def delete(self, request, pk):
        todo = get_object_or_404(Todo, pk=pk)
        todo.delete()
        return HttpResponse('')

    def put(self, request, pk):
        todo = get_object_or_404(Todo, pk=pk)
        data = parse_qs(request.body.decode())
        new_task = data.get('new-task', [''])[0]
        todo.task = new_task
        todo.save()
        return render(request, 'todo/partials/list-task.html', {'todo': todo})


class Edit(View):
    def get(self, request, pk):
        todo = get_object_or_404(Todo, pk=pk)
        return render(request, 'todo/partials/edit-task.html', {'todo': todo})


class Completed(View):
    def get(self, request, pk):
        todo = get_object_or_404(Todo, pk=pk)
        todo.completed = not todo.completed
        todo.save()
        return render(request, 'todo/partials/list-task.html', {'todo': todo})


class FilterTasks(View):
    def get(self, request, filter_type):
        if filter_type == 'completed':
            todo = Todo.objects.filter(completed=True)
        elif filter_type == 'pending':
            todo = Todo.objects.filter(completed=False)
        elif filter_type == 'all':
            todo = Todo.objects.all()
        else:
            return redirect('todo:home')

        return render(
            request,
            'todo/partials/filter-response.html',
            {'todo': todo, 'filter_type': filter_type},
        )
