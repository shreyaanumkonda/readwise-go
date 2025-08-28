# Test Files for Readwise API

This directory contains **safe test utilities** for the Readwise API project.

## **IMPORTANT: These Files Are SAFE for Public Repositories**

- **All data is fictional** - No real user information
- **Test data only** - Used solely for development and testing
- **No sensitive credentials** - All credentials come from environment variables
- **Educational purpose** - Helps developers understand how to use the API

## **Files in This Directory**

### `run_tests.go`
- **Purpose**: Simple test runner that provides testing instructions
- **What it does**: Shows how to test the API with Postman
- **Safety**: No database operations, just guidance

## **How to Use**

### 1. **Prerequisites**
- MySQL server running (XAMPP recommended)
- Database `highlights` created
- Environment variables set (see `.envrc` in root)

### 2. **Run Tests**
```bash
# From the project root directory
go run tests/run_tests.go
```

### 3. **What Happens**
- Shows testing instructions
- Provides Postman endpoint examples
- Guides you through manual testing process

## **Security Features**

- **Environment Variables**: Database credentials from environment, not hardcoded
- **Fictional Data**: All test data is completely made up
- **No Production Data**: Never touches real user data
- **Isolated Testing**: Creates separate test records

## **Testing Approach**

This simplified test runner provides:
- Clear testing instructions
- Postman endpoint examples
- Manual testing guidance
- No automatic data population (to avoid complexity)

## **Why This Is Safe**

1. **No Real Credentials**: All database info comes from environment
2. **Fictional Content**: Book titles, authors, and text are made up
3. **Test Purpose Only**: Designed for development, not production
4. **Clear Documentation**: Every file explains its purpose
5. **Standard Practice**: This is how professional projects handle testing

## **What NOT to Do**

- Don't use real user data
- Don't hardcode database passwords
- Don't commit `.envrc.local` files
- Don't use production database for testing

## **After Running Tests**

Once tests complete successfully, you can:

1. **Start the main server**: `go run .`
2. **Test with Postman**: Use the endpoints shown in the main README
3. **Verify data**: Check your database for the test records
4. **Clean up**: Optionally delete test data when done

## **Contributing**

When adding new test files:
- Use only fictional data
- Document what the test does
- Include safety notes
- Follow the existing pattern

---

**Remember**: These files are designed to help developers learn and test safely. They contain no sensitive information and follow security best practices.
