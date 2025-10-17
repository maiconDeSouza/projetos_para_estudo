from django.shortcuts import render, get_object_or_404, HttpResponse
from django.views import View

from .models import LikePost
from apps.posts.models import Post


# Create your views here.
class LikedPost(View):
    def post(self, request):
        if not request.user.is_authenticated:
            return HttpResponse(status=401)

        pk = request.POST.get('post-link')
        user = request.user

        post = get_object_or_404(Post, pk=pk)

        like, created = LikePost.objects.get_or_create(post=post, user=user)
        likes = post.likes.count()

        liked = False
        if request.user.is_authenticated:
            liked = LikePost.objects.filter(post=post, user=user).exists()

        context = {
            'post': post,
            'likes': likes,
            'liked': liked,
        }
        if created:
            return render(request, 'engagements/partials/liked.html', context)

        like.delete()
        return render(request, 'engagements/partials/no-liked.html', context)
