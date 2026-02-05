from django.urls import path
from . import views

urlpatterns = [
    path('', views.Index.as_view(), name='index'),
    path('delete/<uuid:pk>/', views.Delete.as_view(), name='delete'),
]
