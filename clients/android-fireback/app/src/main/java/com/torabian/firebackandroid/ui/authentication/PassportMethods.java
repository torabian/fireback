package com.torabian.firebackandroid.ui.authentication;

import android.app.Activity;
import android.content.Context;
import android.os.Bundle;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentActivity;
import androidx.navigation.NavController;
import androidx.navigation.NavHostController;
import androidx.navigation.Navigation;
import androidx.navigation.fragment.NavHostFragment;
import androidx.navigation.ui.NavigationUI;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;
import android.widget.Toast;

import com.torabian.firebackandroid.R;

/**
 * A simple {@link Fragment} subclass.
 * Use the {@link PassportMethods#newInstance} factory method to
 * create an instance of this fragment.
 */
public class PassportMethods extends Fragment {

    // TODO: Rename parameter arguments, choose names that match
    // the fragment initialization parameters, e.g. ARG_ITEM_NUMBER
    private static final String ARG_PARAM1 = "param1";
    private static final String ARG_PARAM2 = "param2";

    // TODO: Rename and change types of parameters
    private String mParam1;
    private String mParam2;

    public PassportMethods() {
        // Required empty public constructor
    }

    /**
     * Use this factory method to create a new instance of
     * this fragment using the provided parameters.
     *
     * @param param1 Parameter 1.
     * @param param2 Parameter 2.
     * @return A new instance of fragment PassportMethods.
     */
    // TODO: Rename and change types and number of parameters
    public static PassportMethods newInstance(String param1, String param2) {
        PassportMethods fragment = new PassportMethods();
        Bundle args = new Bundle();
        args.putString(ARG_PARAM1, param1);
        args.putString(ARG_PARAM2, param2);
        fragment.setArguments(args);
        return fragment;
    }

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        if (getArguments() != null) {
            mParam1 = getArguments().getString(ARG_PARAM1);
            mParam2 = getArguments().getString(ARG_PARAM2);
        }

    }

    private FragmentActivity context;

    @Override
    public void onAttach(@NonNull Activity activity) {
        super.onAttach(activity);
        context = (FragmentActivity) activity;
    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);


        NavController navController = Navigation.findNavController(requireActivity(), R.id.fragmentContainerView);

        view.findViewById(R.id.continue_with_email_btn).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view2) {
                navController.navigate(R.id.action_passportMethods_to_continueWithEmail3);
            }
        });

        view.findViewById(R.id.continue_with_phone_btn).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                navController.navigate(R.id.action_passportMethods_to_continueWithPhone);
            }
        });
    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        // Inflate the layout for this fragment
        return inflater.inflate(R.layout.fragment_passport_methods, container, false);
    }
}