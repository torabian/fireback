package com.torabian.firebackandroid.ui.authentication;

import android.content.Intent;
import android.os.Bundle;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.databinding.DataBindingUtil;
import androidx.fragment.app.Fragment;
import androidx.lifecycle.ViewModelProvider;
import androidx.navigation.NavController;
import androidx.navigation.Navigation;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import com.fireback.SessionManager;
import com.fireback.SingleResponse;
import com.fireback.modules.workspaces.ClassicSigninAction;
import com.fireback.modules.workspaces.PostPassportsSigninClassic;
import com.fireback.modules.workspaces.UserSessionDto;
import com.torabian.firebackandroid.MainActivity;
import com.torabian.firebackandroid.R;
import com.torabian.firebackandroid.databinding.FragmentEnterPasswordBinding;

import io.reactivex.rxjava3.android.schedulers.AndroidSchedulers;
import io.reactivex.rxjava3.observers.DisposableSingleObserver;

/**
 * A simple {@link Fragment} subclass.
 * Use the {@link EnterPassword#newInstance} factory method to
 * create an instance of this fragment.
 */
public class EnterPassword extends Fragment {

    // TODO: Rename parameter arguments, choose names that match
    // the fragment initialization parameters, e.g. ARG_ITEM_NUMBER
    private static final String ARG_VALUE = "param1";

    // TODO: Rename and change types of parameters
    private String mValue;

    private ClassicSigninAction.ReqViewModel viewModel;

    public EnterPassword() {
        // Required empty public constructor
    }

    /**
     * Use this factory method to create a new instance of
     * this fragment using the provided parameters.
     *
     * @param param1 Parameter 1.
     * @param param2 Parameter 2.
     * @return A new instance of fragment EnterPassword.
     */
    // TODO: Rename and change types and number of parameters
    public static EnterPassword newInstance(String value) {
        EnterPassword fragment = new EnterPassword();
        Bundle args = new Bundle();
        args.putString(ARG_VALUE, value);
        fragment.setArguments(args);
        return fragment;
    }

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        if (getArguments() != null) {
            mValue = getArguments().getString(ARG_VALUE);
        }
    }

    public void submit() {
        NavController navController = Navigation.findNavController(requireActivity(), R.id.fragmentContainerView);

        PostPassportsSigninClassic action = new PostPassportsSigninClassic();
        ClassicSigninAction.Req dto = new ClassicSigninAction.Req();
        dto.value = viewModel.getValue().getValue();
        dto.password = viewModel.getPassword().getValue();
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

            }
        });

    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);
        view.findViewById(R.id.continue_btn).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                submit();
            }
        });
    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        // Inflate the layout for this fragment
        FragmentEnterPasswordBinding binding = DataBindingUtil.inflate(inflater, R.layout.fragment_enter_password, container, false);
        viewModel = new ViewModelProvider(this).get(ClassicSigninAction.ReqViewModel.class);
        viewModel.setValue(getArguments().getString("value"));
        binding.setViewModel(viewModel);
        binding.setLifecycleOwner(getViewLifecycleOwner());
        return binding.getRoot();
    }
}