from django.shortcuts import HttpResponse
from django.views import View


# Create your views here.
class PostList(View):
    def get(self, request):
        return HttpResponse('Post')
