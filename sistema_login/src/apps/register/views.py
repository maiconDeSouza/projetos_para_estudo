from django.shortcuts import render, redirect
from django.views import View
from django.contrib import messages
from django.contrib.auth.hashers import check_password

from .models import UserLogin

from .forms import UserLoginForm, RegisterForm


# Create your views here.
class Login(View):
    def get(self, request):
        forms = UserLoginForm()
        return render(request, 'register/pages/login.html', {'forms': forms})

    def post(self, request):
        forms = UserLoginForm(request.POST)
        if forms.is_valid():
            email = forms.cleaned_data['email']
            raw = forms.cleaned_data['password']

            try:
                user = UserLogin.objects.get(email=email)
                if check_password(raw, user.password):
                    request.session['user_id'] = user.id
                    return redirect('register:home')
                else:
                    messages.error(request, 'Senha incorreta.')
            except UserLogin.DoesNotExist:
                messages.error(request, 'Usuário não encontrado.')


class Register(View):
    def get(self, request):
        forms = RegisterForm()
        return render(request, 'register/pages/register.html', {'forms': forms})

    def post(self, request):
        forms = RegisterForm(request.POST)

        if forms.is_valid():
            forms.save()
            messages.success(request, 'Cadastro realizado com sucesso!')
            return redirect('register:login')
        else:
            forms = RegisterForm()
            return render(request, 'cadastro.html', {'forms': forms})


class Home(View):
    def get(self, request):
        user_id = request.session.get('user_id')
        user = UserLogin.objects.get(pk=user_id)
        return render(request, 'register/pages/home.html', {'user': user})
