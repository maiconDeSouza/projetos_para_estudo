from django.urls import path
from django.contrib.auth import views as auth_views


from . import views

app_name = 'accounts'
urlpatterns = [
    path('singnup/', views.Signup.as_view(), name='signup'),
    path('logout/', auth_views.LogoutView.as_view(), name='logout'),
]
