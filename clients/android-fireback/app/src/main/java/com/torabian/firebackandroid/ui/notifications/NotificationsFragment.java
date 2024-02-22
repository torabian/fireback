package com.torabian.firebackandroid.ui.notifications;

import android.content.Intent;
import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.fragment.app.Fragment;
import androidx.lifecycle.ViewModelProvider;

import com.fireback.SessionManager;
import com.torabian.firebackandroid.MainActivity;
import com.torabian.firebackandroid.databinding.FragmentNotificationsBinding;

public class NotificationsFragment extends Fragment {

    private FragmentNotificationsBinding binding;

    public View onCreateView(@NonNull LayoutInflater inflater,
                             ViewGroup container, Bundle savedInstanceState) {
        NotificationsViewModel notificationsViewModel =
                new ViewModelProvider(this).get(NotificationsViewModel.class);

        binding = FragmentNotificationsBinding.inflate(inflater, container, false);
        View root = binding.getRoot();

        binding.button2.setOnClickListener(view -> {
            Toast.makeText(getContext().getApplicationContext(), "Signout", Toast.LENGTH_LONG).show();
            SessionManager.getInstance(getContext().getApplicationContext()).signout();
            Intent in = new Intent(getContext().getApplicationContext(), MainActivity.class);
            startActivity(in);
            getActivity().finishAffinity();

     });

        final TextView textView = binding.textNotifications;
        notificationsViewModel.getText().observe(getViewLifecycleOwner(), textView::setText);
        return root;
    }

    @Override
    public void onDestroyView() {
        super.onDestroyView();
        binding = null;
    }
}