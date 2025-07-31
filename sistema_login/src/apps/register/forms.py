from django import forms
from django.contrib.auth.hashers import make_password

from .models import UserLogin


class UserLoginForm(forms.ModelForm):
    email = forms.EmailField(
        label='email',
        widget=forms.EmailInput(attrs={'placeholder': 'Digite seu e-mail'}),
    )
    password = forms.CharField(
        label='Senha',
        widget=forms.PasswordInput(attrs={'placeholder': 'Digite sua senha'}),
    )

    class Meta:
        model = UserLogin
        fields = ['email', 'password']


class RegisterForm(forms.ModelForm):
    password1 = forms.CharField(label='Senha', widget=forms.PasswordInput())
    password2 = forms.CharField(
        label='Confirmação de senha', widget=forms.PasswordInput()
    )

    class Meta:
        model = UserLogin
        fields = ['name', 'email']

    def clean_email(self):
        email = self.cleaned_data['email']
        if not email and UserLogin.objects.filter(email=email).exists():
            raise forms.ValidationError('Este e‑mail já está em uso.')
        return email

    def clean(self):
        cleaned = super().clean()
        p1 = cleaned.get('password1')
        p2 = cleaned.get('password2')

        if p1 and p2 and p1 != p2:
            raise forms.ValidationError('As senhas não coincidem.')
        return cleaned

    def save(self, commit=True):
        user = super().save(commit=False)
        raw = self.cleaned_data['password1']
        user.password = make_password(raw)

        if commit:
            user.save()
            return user
