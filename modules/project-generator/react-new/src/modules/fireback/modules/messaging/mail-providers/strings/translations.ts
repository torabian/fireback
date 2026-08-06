/**
* Auto generated file by fireback language & translation manager.
*/
export const en = {
  emailProviders: {
    fieldApiKey: "API Key",
    fieldCurlScript: "Curl script",
    fieldCurlScriptDescription: "Curl script, which would be called upon sending email.\n          Kindly beaware, this is semantic rather\n          than actual bash script, so use limited features and no extra bash calls.\n          \u003cbr /\u003e\n          Make sure, you put the secrets, templates, credentials, here, there will be no\n          other place to use. While sending an email, following variables will be\n          replaced in your curl message.\n          \u003cbr /\u003e\n\n          %FromName%  string \u003cbr /\u003e\n          %FromEmail% string \u003cbr /\u003e\n          %ToName%    string \u003cbr /\u003e\n          %ToEmail%   string \u003cbr /\u003e\n          %Subject%   string \u003cbr /\u003e\n          %Content%   string (It will be escaped with \\\", so use double qoutes) \u003cbr /\u003e",
    fieldDomain: "Domain",
    fieldHost: "Host",
    fieldPassword: "Password",
    fieldPort: "Port",
    fieldServerToken: "Server Token",
    fieldUsername: "Username",
    mailSent: "Mail has been sent",
    sendTestEmail: "Send a test email",
    title: "Title",
    titleHint: "Title of the email provider, to search and allocate easier.",
    typeCurl: "Curl",
    typeMailgun: "Mailgun",
    typePostmark: "Postmark",
    typeResend: "Resend",
    typeSendgrid: "SendGrid",
    typeSmtp: "SMTP",
    typeTerminal: "Terminal (Debug)",
  },
};/**
* Auto generated file by fireback language & translation manager.
*/
export const fa = {
  emailProviders: {
    fieldApiKey: "کلید API",
    fieldCurlScript: "اسکریپت Curl",
    fieldCurlScriptDescription: "اسکریپت Curl که هنگام ارسال ایمیل فراخوانی می‌شود.\n          لطفاً توجه داشته باشید که این معنایی است تا\n          واقعاً اسکریپت بش، پس محدود به همین امکانات باشید و فراخوانی اضافی بش نداشته باشید.\n          \u003cbr /\u003e\n          مطمئن شوید اسرار، قالب‌ها و اعتبارنامه‌ها را همین‌جا قرار می‌دهید، جای دیگری برای\n          استفاده از آن‌ها وجود نخواهد داشت. هنگام ارسال ایمیل، متغیرهای زیر در\n          پیام curl شما جایگزین خواهند شد.\n          \u003cbr /\u003e\n\n          %FromName%  رشته \u003cbr /\u003e\n          %FromEmail% رشته \u003cbr /\u003e\n          %ToName%    رشته \u003cbr /\u003e\n          %ToEmail%   رشته \u003cbr /\u003e\n          %Subject%   رشته \u003cbr /\u003e\n          %Content%   رشته (با \\\" اسکیپ خواهد شد، پس از دابل‌کوتیشن استفاده کنید) \u003cbr /\u003e",
    fieldDomain: "دامنه",
    fieldHost: "میزبان",
    fieldPassword: "رمز عبور",
    fieldPort: "پورت",
    fieldServerToken: "توکن سرور",
    fieldUsername: "نام کاربری",
    mailSent: "ایمیل ارسال شد",
    sendTestEmail: "ارسال ایمیل آزمایشی",
    title: "عنوان",
    titleHint: "عنوان ارائه‌دهنده ایمیل، برای جستجو و تخصیص آسان‌تر.",
    typeCurl: "Curl",
    typeMailgun: "Mailgun",
    typePostmark: "Postmark",
    typeResend: "Resend",
    typeSendgrid: "SendGrid",
    typeSmtp: "SMTP",
    typeTerminal: "ترمینال (دیباگ)",
  },
};/**
* Auto generated file by fireback language & translation manager.
*/
export const pl = {
  emailProviders: {
    fieldApiKey: "Klucz API",
    fieldCurlScript: "Skrypt Curl",
    fieldCurlScriptDescription: "Skrypt Curl, który zostanie wywołany podczas wysyłania wiadomości e-mail.\n          Należy pamiętać, że jest to podejście semantyczne, a nie\n          rzeczywisty skrypt bash, więc używaj ograniczonych funkcji i bez dodatkowych wywołań bash.\n          \u003cbr /\u003e\n          Upewnij się, że umieszczasz tutaj sekrety, szablony i dane uwierzytelniające, nie będzie\n          innego miejsca do ich użycia. Podczas wysyłania wiadomości e-mail poniższe zmienne zostaną\n          zastąpione w Twojej wiadomości curl.\n          \u003cbr /\u003e\n\n          %FromName%  ciąg znaków \u003cbr /\u003e\n          %FromEmail% ciąg znaków \u003cbr /\u003e\n          %ToName%    ciąg znaków \u003cbr /\u003e\n          %ToEmail%   ciąg znaków \u003cbr /\u003e\n          %Subject%   ciąg znaków \u003cbr /\u003e\n          %Content%   ciąg znaków (zostanie zabezpieczony znakiem \\\", więc używaj cudzysłowów podwójnych) \u003cbr /\u003e",
    fieldDomain: "Domena",
    fieldHost: "Host",
    fieldPassword: "Hasło",
    fieldPort: "Port",
    fieldServerToken: "Token serwera",
    fieldUsername: "Nazwa użytkownika",
    mailSent: "Wiadomość e-mail została wysłana",
    sendTestEmail: "Wyślij testową wiadomość e-mail",
    title: "Tytuł",
    titleHint: "Tytuł dostawcy poczty e-mail, aby ułatwić wyszukiwanie i przypisywanie.",
    typeCurl: "Curl",
    typeMailgun: "Mailgun",
    typePostmark: "Postmark",
    typeResend: "Resend",
    typeSendgrid: "SendGrid",
    typeSmtp: "SMTP",
    typeTerminal: "Terminal (debugowanie)",
  },
};
 export const strings = {...en, $fa:fa,$pl:pl};
