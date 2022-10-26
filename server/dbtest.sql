CREATE user:dd81b0581396472a97b9ded2813a6cc3 SET 
    email = 'test@gmail.com',  
    email_verified_at = 1666768660,
    pasword = 'test2004',
    sign_up_method = 'github',
    given_name = 'test',
    first_name = 'first',
    middle_name = 'middle',
    last_name = 'last',
    full_name = string::join(' ', first_name, middle_name, last_name),
    nick_name = 'tester',
    gender = 'male',
    birth_date = '2022.04.05',
    phone_number = '0773456777693',
    phone_number_verified_at = 1666768660,
    picture = 'minio object store url',
    roles = ['user', 'admin', 'manager', 'maaintainer'],
    default_role = roles[0],
    assigned_role = roles[1],
    assigned_roles = [default_role, assigned_role],
    revoked_timestamp = 1666768660,
    is_multi_factor_auth_enabled = true,
    updated_at = 1666768660,
    created_at = 1666768660,
    last_logged_in = 1666768660
;