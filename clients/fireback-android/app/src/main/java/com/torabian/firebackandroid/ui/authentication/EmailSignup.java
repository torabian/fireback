package com.torabian.firebackandroid.ui.authentication;

import androidx.databinding.DataBindingUtil;
import androidx.lifecycle.ViewModelProvider;

import android.content.Intent;
import android.os.Bundle;

import androidx.fragment.app.Fragment;
import androidx.navigation.NavController;
import androidx.navigation.Navigation;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import com.fireback.SessionManager;
import com.fireback.SingleResponse;
import com.fireback.modules.workspaces.ClassicAuthDto;
import com.fireback.modules.workspaces.ClassicSigninAction;
import com.fireback.modules.workspaces.ClassicSignupAction;
import com.fireback.modules.workspaces.PostPassportsSigninClassic;
import com.fireback.modules.workspaces.PostPassportsSignupClassic;
import com.fireback.modules.workspaces.UserSessionDto;
import com.torabian.firebackandroid.MainActivity;
import com.torabian.firebackandroid.R;
import com.torabian.firebackandroid.databinding.FragmentEmailSignupBinding;
import com.torabian.firebackandroid.databinding.FragmentEnterPasswordBinding;

import io.reactivex.rxjava3.android.schedulers.AndroidSchedulers;
import io.reactivex.rxjava3.annotations.NonNull;
import io.reactivex.rxjava3.annotations.Nullable;
import io.reactivex.rxjava3.observers.DisposableSingleObserver;

public class EmailSignup extends Fragment {

    private ClassicAuthDto.VM mViewModel;

    public static EmailSignup newInstance() {
        return new EmailSignup();
    }

    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container,
                             @Nullable Bundle savedInstanceState) {

        FragmentEmailSignupBinding binding = DataBindingUtil.inflate(inflater, R.layout.fragment_email_signup, container, false);
        mViewModel = new ViewModelProvider(this).get(ClassicAuthDto.VM.class);
        mViewModel.setValue(getArguments().getString("value"));
        binding.setViewModel(mViewModel);
        binding.setLifecycleOwner(getViewLifecycleOwner());
        return binding.getRoot();
    }

    @Override
    public void onActivityCreated(@Nullable Bundle savedInstanceState) {
        super.onActivityCreated(savedInstanceState);
        mViewModel = new ViewModelProvider(this).get(ClassicAuthDto.VM.class);
        // TODO: Use the ViewModel
    }

    @Override
    public void onViewCreated(@androidx.annotation.NonNull View view, @androidx.annotation.Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);


        view.findViewById(R.id.email_signup_continue_btn).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                submit();
            }
        });
    }


    public void submit() {

        PostPassportsSignupClassic action = new PostPassportsSignupClassic();
        ClassicSignupAction.Req dto = new ClassicSignupAction.Req();
        Toast.makeText(getContext().getApplicationContext(), "1", Toast.LENGTH_SHORT).show();

        dto.value = mViewModel.getValue().getValue();
        dto.password = mViewModel.getPassword().getValue();
        dto.firstName = mViewModel.getFirstName().getValue();
        dto.lastName = mViewModel.getLastName().getValue();
        dto.type = "email";
        Toast.makeText(getActivity(), "Begin", Toast.LENGTH_SHORT).show();


        action.post(dto).observeOn(AndroidSchedulers.mainThread()).subscribe(new DisposableSingleObserver<SingleResponse<UserSessionDto>>() {
            @Override
            public void onSuccess(@io.reactivex.rxjava3.annotations.NonNull SingleResponse<UserSessionDto> userSessionDtoSingleResponse) {
                if (userSessionDtoSingleResponse != null) {
                    SessionManager.getInstance(getActivity().getApplicationContext()).saveUserSession(userSessionDtoSingleResponse.data);
                    Toast.makeText(getActivity(), "Done", Toast.LENGTH_SHORT).show();
                    Intent in = new Intent(getContext().getApplicationContext(), MainActivity.class);
                    startActivity(in);
                    getActivity().finishAffinity();
                } else {
                    Toast.makeText(getActivity(), "Fail", Toast.LENGTH_SHORT).show();
                }
            }

            @Override
            public void onError(@io.reactivex.rxjava3.annotations.NonNull Throwable e) {
                Toast.makeText(getActivity(), "Error" + e.toString(), Toast.LENGTH_SHORT).show();

            }
        });

    }
}