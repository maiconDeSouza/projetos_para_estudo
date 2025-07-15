from django.urls import path
from django.contrib.auth import views as auth_views

from .views import signup, logout

urlpatterns = [
    path('signup/', signup, name='signup'),
    path(
        'login/',
        auth_views.LoginView.as_view(template_name='accounts/login.html'),
        name='login'
    ),
    path(
        'logged_out/',
        auth_views.LogoutView.as_view(
            template_name='accounts/logged_out.html'),
        name='logged_out'
    ),
    path('logout/', logout, name='logout'),
]
