from django.shortcuts import render, redirect
from django.views import View
from django.contrib import messages

from .forms import SignupForm


# Create your views here.
class Signup(View):
    template_name = 'accounts/pages/signup.html'
    form_class = SignupForm

    def get(self, request):
        form = self.form_class()
        return render(request, self.template_name, {'form': form})

    def post(self, request):
        form = self.form_class(request.POST)
        if form.is_valid():
            form.save()
            return redirect('accounts:login')
        messages.error(request, 'Por favor corrija os erros abaixo.')
        return render(request, self.template_name, {'form': form})
