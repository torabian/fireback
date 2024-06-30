package com.torabian.firebackandroid.ui.workspaces;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;
import androidx.swiperefreshlayout.widget.SwipeRefreshLayout;

import com.fireback.ArrayResponse;
import com.fireback.modules.workspaces.GetUsers;
import com.fireback.modules.workspaces.PersonEntity;
import com.fireback.modules.workspaces.UserEntity;
import com.torabian.firebackandroid.R;

import java.util.ArrayList;
import java.util.List;

import io.reactivex.rxjava3.android.schedulers.AndroidSchedulers;
import io.reactivex.rxjava3.core.Observer;
import io.reactivex.rxjava3.disposables.Disposable;

/**
 * A simple {@link Fragment} subclass.
 * Use the {@link UserList#newInstance} factory method to
 * create an instance of this fragment.
 */
public class UserList extends Fragment {
    SwipeRefreshLayout r;
    // TODO: Rename parameter arguments, choose names that match
    // the fragment initialization parameters, e.g. ARG_ITEM_NUMBER
    private static final String ARG_PARAM1 = "param1";
    private static final String ARG_PARAM2 = "param2";

    // TODO: Rename and change types of parameters
    private String mParam1;
    private String mParam2;

    public UserList() {
        // Required empty public constructor
    }

    /**
     * Use this factory method to create a new instance of
     * this fragment using the provided parameters.
     *
     * @param param1 Parameter 1.
     * @param param2 Parameter 2.
     * @return A new instance of fragment RoleList.
     */
    // TODO: Rename and change types and number of parameters
    public static UserList newInstance(String param1, String param2) {
        UserList fragment = new UserList();
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

    private List<UserEntity> computedItems;

    private List<UserEntity> getRoles() {


        List<UserEntity> roles = new ArrayList<>();

        for (int i = 1; i <= 20; i++) {
            UserEntity r1 = new UserEntity();
            r1.person = new PersonEntity();
            r1.person.firstName = "Ali";
            r1.person.lastName = "Torabi";
            roles.add(r1);
        }

        return roles;
    }

    private void getRolesRx() {
        GetUsers roles = new GetUsers(getContext());
        roles.query()
                .observeOn(AndroidSchedulers.mainThread())
                .subscribe(new Observer<ArrayResponse<UserEntity>>() {
                    @Override
                    public void onSubscribe(@io.reactivex.rxjava3.annotations.NonNull Disposable d) {

                    }

                    @Override
                    public void onNext(@io.reactivex.rxjava3.annotations.NonNull ArrayResponse<UserEntity> userEntityArrayResponse) {
                        System.out.println(userEntityArrayResponse.data.items);
                        adapter.setCardItems(userEntityArrayResponse.data.items);
                        r.setRefreshing(false);
                    }

                    @Override
                    public void onError(@io.reactivex.rxjava3.annotations.NonNull Throwable e) {
                        Toast.makeText(getActivity(), "Error:" + e.toString(), Toast.LENGTH_SHORT).show();
                        r.setRefreshing(false);
                    }

                    @Override
                    public void onComplete() {
                        r.setRefreshing(false);
                    }
                });


    }
    UserAdapter adapter;

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {



        super.onViewCreated(view, savedInstanceState);
        adapter = new UserAdapter();
        adapter.setCardItems(new ArrayList<>());


        view.findViewById(R.id.floating_btn).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
            

            }
        });


        RecyclerView v = view.findViewById(R.id.role_list_recycler);
        r = view.findViewById(R.id.role_list_refresh);
        r.setOnRefreshListener(new SwipeRefreshLayout.OnRefreshListener() {
            @Override
            public void onRefresh() {
                getRolesRx();

            }
        });

        v.setLayoutManager(new LinearLayoutManager(getContext()));
        v.setAdapter(adapter);

        getRolesRx();

    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        // Inflate the layout for this fragment
        return inflater.inflate(R.layout.fragment_role_list, container, false);
    }
}