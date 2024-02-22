package com.torabian.firebackandroid.ui.home;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import io.reactivex.rxjava3.android.schedulers.AndroidSchedulers;
import io.reactivex.rxjava3.core.Observable;
import io.reactivex.rxjava3.schedulers.Schedulers;
import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.lifecycle.ViewModelProvider;

import com.fireback.modules.workspaces.PersonEntity;
import com.fireback.modules.workspaces.UserEntity;
import com.torabian.firebackandroid.R;
import com.torabian.firebackandroid.databinding.FragmentHomeBinding;

public class HomeFragment extends Fragment {

    private FragmentHomeBinding binding;

    public View onCreateView(@NonNull LayoutInflater inflater,
                             ViewGroup container, Bundle savedInstanceState) {
        HomeViewModel homeViewModel =
                new ViewModelProvider(this).get(HomeViewModel.class);

        binding = FragmentHomeBinding.inflate(inflater, container, false);
        View root = binding.getRoot();


        System.out.println("This is the home!! ");
//
//        Observable.range(1, 5).
//                subscribeOn(Schedulers.io())
//                .observeOn(AndroidSchedulers.mainThread())
//                .subscribe(value -> System.out.println("Event happened:" + "" + value));
//        final TextView textView = binding.textHome;

         return root;
    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);

//        TextView txt = view.findViewById(R.id.text_home);
//        txt.setText("Loading...");

        // Create a new ObjectAnimator for the TextView's alpha property
//        ObjectAnimator fadeIn = ObjectAnimator.ofFloat(txt, "alpha", 0f, 1f);
//        fadeIn.setDuration(5000); // Set the duration of the animation in milliseconds (e.g., 2000ms or 2 seconds)
//        txt.setVisibility(View.VISIBLE); // Ensure the TextView is visible before starting the animation
//        fadeIn.start(); // Start the fade-in animation
//
//        Observable.range(1,5).
//                concatMap(i -> Observable.just(i).delay(300, TimeUnit.MILLISECONDS)).
//                subscribeOn(Schedulers.io()).
//                observeOn(AndroidSchedulers.mainThread()).
//                subscribe(value -> txt.setText("Counting" + value));


        UserEntity user = new UserEntity();
        user.person = new PersonEntity();
        user.person.firstName = "Ali";
        user.person.lastName = "Torabi";

        System.out.println("Getting session:");


     }

    @Override
    public void onDestroyView() {
        super.onDestroyView();
        binding = null;
    }
}