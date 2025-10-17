from django.shortcuts import render, get_object_or_404
from django.views import View
from django.views.generic import ListView

from .models import Post
from apps.engagements.models import LikePost


# Create your views here.
# class Home(View):
#     def get(self, request):
#         posts = Post.objects.filter(published=True)

#         context = {
#             'posts': posts,
#         }

#         return render(request, 'posts/pages/index.html', context)


class Home(ListView):
    model = Post
    template_name = 'posts/pages/index.html'
    context_object_name = 'posts'
    paginate_by = 9

    def get_queryset(self):
        return Post.objects.filter(published=True)


class PostDetail(View):
    def get(self, request, slug):
        user = request.user
        post = get_object_or_404(Post, slug=slug)

        likes = post.likes.count()

        liked = False
        if request.user.is_authenticated:
            liked = LikePost.objects.filter(post=post, user=user).exists()

        context = {
            'post': post,
            'likes': likes,
            'liked': liked,
        }

        return render(request, 'posts/pages/details.html', context)
