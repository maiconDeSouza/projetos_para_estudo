from django.shortcuts import render
from django.views import View

from .models import Post


# Create your views here.
class Home(View):
    def get(self, request):
        posts = Post.objects.filter(published=True)

        context = {
            'posts': posts,
        }

        return render(request, 'posts/pages/index.html', context)
