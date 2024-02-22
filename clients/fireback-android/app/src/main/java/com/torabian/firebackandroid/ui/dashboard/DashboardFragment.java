package com.torabian.firebackandroid.ui.dashboard;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.lifecycle.ViewModelProvider;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.torabian.firebackandroid.R;
import com.torabian.firebackandroid.databinding.FragmentDashboardBinding;
import com.torabian.firebackandroid.ui.CustomAdapter;

import java.util.ArrayList;
import java.util.List;

public class DashboardFragment extends Fragment {

    private FragmentDashboardBinding binding;

    public View onCreateView(@NonNull LayoutInflater inflater,
                             ViewGroup container, Bundle savedInstanceState) {
        DashboardViewModel dashboardViewModel =
                new ViewModelProvider(this).get(DashboardViewModel.class);

        binding = FragmentDashboardBinding.inflate(inflater, container, false);
        View root = binding.getRoot();


        return root;
    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);
//
//        RecyclerView v =  view.findViewById(R.id.recycler2);
//        CustomAdapter adapter = new CustomAdapter();
//        List<String> items = new ArrayList<>();
//        for (int i = 0 ; i <= 1000; i ++ ) {
//            items.add("Item number: " + i);
//        }
//
//        adapter.setCardItems(items);
//        v.setLayoutManager(new LinearLayoutManager(getContext()));
//        v.setAdapter(adapter);
    }

    @Override
    public void onDestroyView() {
        super.onDestroyView();
        binding = null;
    }
}