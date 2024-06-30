package com.torabian.firebackandroid;

import android.content.Intent;
import android.hardware.camera2.CameraExtensionSession;
import android.os.Bundle;

import com.fireback.FirebackConfig;
import com.fireback.SessionManager;
import com.fireback.modules.workspaces.UserSessionDto;
import com.google.android.material.bottomnavigation.BottomNavigationView;
import androidx.activity.result.contract.ActivityResultContracts;
import androidx.appcompat.app.AppCompatActivity;
import androidx.camera.core.CameraX;
import androidx.camera.core.ImageCapture;
import androidx.navigation.NavController;
import androidx.navigation.Navigation;
import androidx.navigation.ui.AppBarConfiguration;
import androidx.navigation.ui.NavigationUI;

import com.torabian.firebackandroid.databinding.ActivityMainBinding;
import com.torabian.firebackandroid.ui.authentication.AuthenticationActivity;
public class MainActivity extends AppCompatActivity {

    private ActivityMainBinding binding;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        FirebackConfig.getInstance().setRemoteUrl(getResources().getString(R.string.api_url));

        UserSessionDto dto = SessionManager.getInstance(getApplicationContext()).getUserSession();

        if (dto == null) {
            Intent in = new Intent(getApplicationContext(), AuthenticationActivity.class);
            startActivity(in);
        } else {
            binding = ActivityMainBinding.inflate(getLayoutInflater());
            setContentView(binding.getRoot());

            BottomNavigationView navView = findViewById(R.id.nav_view);
            // Passing each menu ID as a set of Ids because each


            // menu should be considered as top level destinations.
            AppBarConfiguration appBarConfiguration = new AppBarConfiguration.Builder(
                    R.id.navigation_home, R.id.navigation_roles, R.id.navigation_profile)
                    .build();
            NavController navController = Navigation.findNavController(this, R.id.nav_host_fragment_activity_main);
//            NavigationUI.setupActionBarWithNavController(this, navController, appBarConfiguration);
            NavigationUI.setupWithNavController(binding.navView, navController);

        }
    }

}