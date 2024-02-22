package com.torabian.firebackandroid.ui.authentication;

import androidx.lifecycle.ViewModelProvider;

import android.os.Bundle;

import androidx.fragment.app.Fragment;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import com.fireback.modules.workspaces.ClassicAuthDto;
import com.torabian.firebackandroid.R;

import io.reactivex.rxjava3.annotations.NonNull;
import io.reactivex.rxjava3.annotations.Nullable;

public class EmailSignup extends Fragment {

    private ClassicAuthDto.VM mViewModel;

    public static EmailSignup newInstance() {
        return new EmailSignup();
    }

    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container,
                             @Nullable Bundle savedInstanceState) {
        return inflater.inflate(R.layout.fragment_email_signup, container, false);
    }

    @Override
    public void onActivityCreated(@Nullable Bundle savedInstanceState) {
        super.onActivityCreated(savedInstanceState);
        mViewModel = new ViewModelProvider(this).get(ClassicAuthDto.VM.class);
        // TODO: Use the ViewModel
    }

}