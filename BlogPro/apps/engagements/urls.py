from django.urls import path

from . import views

app_name = 'engagements'
urlpatterns = [
    path('liked/', views.LikedPost.as_view(), name='liked'),
]
