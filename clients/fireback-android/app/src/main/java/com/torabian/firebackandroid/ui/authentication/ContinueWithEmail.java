package com.torabian.firebackandroid.ui.authentication;

import androidx.databinding.DataBindingUtil;
import androidx.lifecycle.Observer;
import androidx.lifecycle.ViewModelProvider;

import android.os.Bundle;

import androidx.fragment.app.Fragment;
import androidx.navigation.NavController;
import androidx.navigation.Navigation;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import com.fireback.IResponseError;
import com.fireback.ResponseErrorException;
import com.fireback.SingleResponse;
import com.fireback.modules.workspaces.CheckClassicPassportAction;
import com.fireback.modules.workspaces.PostWorkspacePassportCheck;
import com.fireback.modules.workspaces.UserSessionDto;
import com.torabian.firebackandroid.R;
import com.torabian.firebackandroid.databinding.FragmentContinueWithEmailBinding;
import com.torabian.firebackandroid.ui.AsyncButton;

import io.reactivex.rxjava3.android.schedulers.AndroidSchedulers;
import io.reactivex.rxjava3.annotations.NonNull;
import io.reactivex.rxjava3.annotations.Nullable;
import io.reactivex.rxjava3.core.Single;
import io.reactivex.rxjava3.functions.Consumer;
import io.reactivex.rxjava3.observers.DisposableSingleObserver;

public class ContinueWithEmail extends Fragment {

    private CheckClassicPassportAction.ReqViewModel mViewModel;

    public static ContinueWithEmail newInstance() {
        return new ContinueWithEmail();
    }

    @Override
    public View onCreateView(
            @NonNull LayoutInflater inflater,
            @Nullable ViewGroup container,
            @Nullable Bundle savedInstanceState
    ) {
        FragmentContinueWithEmailBinding binding = DataBindingUtil.inflate(inflater, R.layout.fragment_continue_with_email, container, false);

        mViewModel = new ViewModelProvider(this).get(CheckClassicPassportAction.ReqViewModel.class);
        binding.setViewModel(mViewModel);
        binding.setLifecycleOwner(getViewLifecycleOwner());

        return binding.getRoot();
    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);
        AsyncButton<CheckClassicPassportAction.Res> btn = view.findViewById(R.id.continue_btn);

        // This is how you bind the view model to the async button which would handle
        // the actual form, values, error messages, etc...
        btn.setAction(() -> this.getAction(mViewModel).doOnSuccess(this::onSuccess));
    }

    public static void CastErrorToModel( Throwable e, CheckClassicPassportAction.ReqViewModel mViewModel ) {
        ResponseErrorException responseError = (ResponseErrorException) e;
        responseError.error.errors.forEach(item -> {
            if (item.location.equals("value")) {
                mViewModel.setValueMsg(item.messageTranslated);
            }
        });
    }

    public static Single<SingleResponse<CheckClassicPassportAction.Res>> getAction( CheckClassicPassportAction.ReqViewModel mViewModel) {
        PostWorkspacePassportCheck action = new PostWorkspacePassportCheck();
        CheckClassicPassportAction.Req dto = new CheckClassicPassportAction.Req();
        dto.value = mViewModel.getValue().getValue();

        return action.post(dto).observeOn(
                        AndroidSchedulers.mainThread()
                )
                .doOnError(e -> {
                    ContinueWithEmail.CastErrorToModel(e, mViewModel);
                });
    }

    public void onSuccess(@io.reactivex.rxjava3.annotations.NonNull SingleResponse<CheckClassicPassportAction.Res> resSingleResponse) {
        NavController navController = Navigation.findNavController(requireActivity(), R.id.fragmentContainerView);

        if (resSingleResponse == null) {
            return;
        }
        Bundle bundle = new Bundle();
        bundle.putString("value", mViewModel.getValue().getValue());

        if (resSingleResponse.data.exists) {
            Toast.makeText(getActivity(), "Account exists", Toast.LENGTH_SHORT).show();
            navController.navigate(R.id.action_continueWithEmail3_to_enterPassword, bundle);
        } else {
            Toast.makeText(getActivity(), "Not exists.", Toast.LENGTH_SHORT).show();
            navController.navigate(R.id.action_continueWithEmail3_to_emailSignup, bundle);
        }
    }
}


