package course

import (
	"errors"

	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/course/course_repository"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/user/user_repository"
)

func (svc *courseService) InjectCourseRepository(repo course_repository.CourseRepository) error {
	if repo != nil {
		svc.courseRepository = repo
		return nil
	}
	return errors.New("course repository not found")
}

func (svc *courseService) InjectUserRepository(repo user_repository.UserRepository) error {
	if repo != nil {
		svc.userRepository = repo
		return nil
	}
	return errors.New("user repository not found")
}
