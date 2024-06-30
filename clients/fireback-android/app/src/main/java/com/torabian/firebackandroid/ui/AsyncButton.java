package com.torabian.firebackandroid.ui;
import android.content.Context;
import android.util.AttributeSet;
import android.view.LayoutInflater;
import android.view.View;
import android.widget.Button;
import android.widget.ProgressBar;
import android.widget.RelativeLayout;
import android.widget.TextView;

import com.fireback.ResponseErrorException;
import com.fireback.SingleResponse;
import com.torabian.firebackandroid.R;

import io.reactivex.rxjava3.annotations.NonNull;
import io.reactivex.rxjava3.core.Single;
import io.reactivex.rxjava3.core.SingleObserver;
import io.reactivex.rxjava3.disposables.Disposable;

public class AsyncButton<T> extends RelativeLayout {

    SingleFunction<SingleResponse<T>> action;
    private Button button;
    private ProgressBar loader;
    private TextView extraText;

    public AsyncButton(Context context) {
        super(context);
        init(context);
    }

    public AsyncButton(Context context, AttributeSet attrs) {
        super(context, attrs);
        init(context);
    }

    public AsyncButton(Context context, AttributeSet attrs, int defStyleAttr) {
        super(context, attrs, defStyleAttr);
        init(context);
    }

    private void init(Context context) {
        LayoutInflater.from(context).inflate(R.layout.async_button_layout, this, true);
        button = findViewById(R.id.async_button_super);
        extraText = findViewById(R.id.async_button_message);
        button.setOnClickListener(new OnClickListener() {
            @Override
            public void onClick(View view) {
                loader.setVisibility(VISIBLE);
                extraText.setText("");
                action.call().subscribe(new SingleObserver<SingleResponse<T>>() {
                    @Override
                    public void onSubscribe(@NonNull Disposable d) {

                    }

                    @Override
                    public void onSuccess(@NonNull SingleResponse<T> tSingleResponse) {
                        loader.setVisibility(GONE);

                    }

                    @Override
                    public void onError(@NonNull Throwable e) {
                        if (e instanceof ResponseErrorException) {
                            ResponseErrorException responseError = (ResponseErrorException) e;
                            extraText.setText(responseError.error.messageTranslated);
                        } else {
                            extraText.setText(e.getLocalizedMessage());
                        }
                        loader.setVisibility(GONE);
                    }

                    @Override
                    protected void finalize() throws Throwable {
                        super.finalize();
                        loader.setVisibility(GONE);
                    }
                });
            }
        });
        loader = findViewById(R.id.async_button_progress);
//        extraText = findViewById(R.id.extra_text);
    }

    public void setAction( SingleFunction<SingleResponse<T>>  res) {
        action = res;
    }

    public void showLoader() {
        loader.setVisibility(VISIBLE);
        button.setVisibility(GONE);
    }

    public void hideLoader() {
        loader.setVisibility(GONE);
        button.setVisibility(VISIBLE);
    }

    public void setExtraText(String text) {
        extraText.setText(text);
        extraText.setVisibility(VISIBLE);
    }
}