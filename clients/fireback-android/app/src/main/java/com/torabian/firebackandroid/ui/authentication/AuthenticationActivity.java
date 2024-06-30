package com.torabian.firebackandroid.ui.authentication;

import androidx.appcompat.app.AppCompatActivity;
import androidx.fragment.app.FragmentManager;
import androidx.fragment.app.FragmentTransaction;

import android.os.Bundle;
import android.view.View;
import android.widget.Button;

import com.fireback.SessionManager;
import com.fireback.SingleResponse;
import com.fireback.modules.workspaces.ClassicSigninAction;
import com.fireback.modules.workspaces.PostPassportsSigninClassic;
import com.fireback.modules.workspaces.UserSessionDto;
import com.torabian.firebackandroid.R;

import io.reactivex.rxjava3.android.schedulers.AndroidSchedulers;
import io.reactivex.rxjava3.observers.DisposableSingleObserver;

public class AuthenticationActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_authentication);

//        Button btn = findViewById(R.id.continueWithMailBtn);
//        btn.setOnClickListener(
//                new View.OnClickListener() {
//                    @Override
//                    public void onClick(View view) {
//
//                        FragmentManager f = getSupportFragmentManager();
//                        FragmentTransaction ts = f.beginTransaction();
//                        ContinueWithEmail fragment = new ContinueWithEmail();
//                        ts.replace(R.id.fragmentContainerView, fragment);
//                        ts.addToBackStack(null);
//                        ts.commit();
//
//                    }
//                }
//        );


//        PostPassportsSigninClassic action = new PostPassportsSigninClassic();
//        ClassicSigninAction.Req dto = new ClassicSigninAction.Req();
//        dto.value = "test";
//        dto.password = "test";
//
//        action.post(dto).observeOn(
//                AndroidSchedulers.mainThread()
//        ).subscribe(new DisposableSingleObserver<SingleResponse<UserSessionDto>>() {
//            @Override
//            public void onSuccess(@io.reactivex.rxjava3.annotations.NonNull SingleResponse<UserSessionDto> userSessionDtoSingleResponse) {
//
//                if (userSessionDtoSingleResponse != null) {
//                    System.out.println(userSessionDtoSingleResponse.data.token);
//                    SessionManager.getInstance(getApplicationContext()).saveUserSession(userSessionDtoSingleResponse.data);
////                    txt.setText("Logged in :)");
//                } else {
////                    txt.setText("Unknown things has happened");
//                }
//            }
//
//            @Override
//            public void onError(@io.reactivex.rxjava3.annotations.NonNull Throwable e) {
//                System.out.println(e);
////                txt.setText("Error:: " + e.toString());
//            }
//        });
    }




}