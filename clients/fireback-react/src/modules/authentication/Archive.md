These functions are coming from fireback itself.

// const mutation = useMutation<LoginFormResponse, unknown, BasicUserAuthForm>(
// content => {
// return PassportActions.fn({
// prefix: 'http://localhost:7000',
// }).postPassportSignupEmail({
// email: content.email,
// password: content.password,
// });
// // return execApi('post', 'auth/user/signup', content);
// },
// );

// const onSubmit = (
// values: BasicUserAuthForm,
// formikProps: FormikHelpers<BasicUserAuthForm>,
// ) => {
// mutation.mutate(values, {
// onSuccess(response) {
// if (response.data) {
// setSession(response.data);
// saveCredentials(values);
// navigation.navigate('app', {screen: Screens.Home});
// }
// },
// onError(error: any) {
// formikProps.setErrors(mutationErrorsToFormik(error));
// },
// });
// };
