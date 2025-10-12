from django.urls import path

from . import views

urlpatterns = [
    path('', views.Home.as_view(), name='home'),
    path('post/<slug:slug>', views.PostDetail.as_view(), name='details'),
]
