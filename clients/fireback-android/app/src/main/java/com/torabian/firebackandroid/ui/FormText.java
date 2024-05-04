package com.torabian.firebackandroid.ui;

import android.content.Context;
import android.text.Editable;
import android.text.TextUtils;
import android.text.TextWatcher;
import android.util.AttributeSet;
import android.view.LayoutInflater;
import android.widget.EditText;
import android.widget.LinearLayout;
import android.widget.TextView;

import androidx.databinding.Bindable;
import androidx.databinding.BindingAdapter;
import androidx.databinding.InverseBindingAdapter;
import androidx.databinding.InverseBindingListener;

import com.torabian.firebackandroid.R;

public class FormText extends LinearLayout {

    private String content3;
    private EditText editText;
    private TextView errorMessage;
    private String errorMessageContent = "";

    public FormText(Context context) {
        super(context);
        init(context);
    }


    @BindingAdapter("android:errorMessage")
    public static void setErrorMessage(FormText view, String text) {
        if (text != null && !text.equals(view.errorMessageContent)) {
            view.errorMessageContent = text;
            view.errorMessage.setText(text);
            System.out.println(text);
        }
    }


    @BindingAdapter("android:text")
    public static void setText(FormText view, String text) {
        if (view.content3 == null || !view.content3.equals(text)) {
            view.editText.setText(text);
            view.content3 = text;
        }
    }

    @InverseBindingAdapter(attribute = "android:text")
    public static String getText(FormText view) {
        return view.editText.getText().toString();
    }

    @BindingAdapter("android:textAttrChanged")
    public static void setTextListener(FormText view, final InverseBindingListener listener) {
        if (listener != null) {
            view.editText.addTextChangedListener(new TextWatcher() {
                @Override
                public void beforeTextChanged(CharSequence s, int start, int count, int after) {
                }

                @Override
                public void onTextChanged(CharSequence s, int start, int before, int count) {
                }

                @Override
                public void afterTextChanged(Editable s) {
                    listener.onChange();
                }
            });
        }
    }

    public FormText(Context context, AttributeSet attrs) {
        super(context, attrs);
        init(context);
    }

    private void init(Context context) {
        LayoutInflater.from(context).inflate(R.layout.form_text_layout, this);

        errorMessage = findViewById(R.id.form_text_error_msg);

        editText = findViewById(R.id.form_text_input);
        editText.addTextChangedListener(new TextWatcher() {
            @Override
            public void beforeTextChanged(CharSequence charSequence, int i, int i1, int i2) {}

            @Override
            public void onTextChanged(CharSequence charSequence, int i, int i1, int i2) {
                content3 = charSequence.toString();
            }

            @Override
            public void afterTextChanged(Editable editable) {}
        });
    }


    public String getText() {
        return editText.getText().toString();
    }

}