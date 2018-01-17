<?php
declare(strict_types=1);

namespace App\Command;

use App\Entity\User;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;
use Symfony\Component\Security\Core\Encoder\UserPasswordEncoderInterface;

class CreateUserCommand  extends Command
{
    /**
     * @var EntityManagerInterface
     */
    private $entityManager;

    /**
     * @var UserPasswordEncoderInterface
     */
    private $encoder;

    /**
     * CreateUserCommand constructor.
     * @param EntityManagerInterface $entityManager
     * @param UserPasswordEncoderInterface $encoder
     */
    public function __construct(EntityManagerInterface $entityManager, UserPasswordEncoderInterface $encoder)
    {
        parent::__construct();
        $this->entityManager = $entityManager;
        $this->encoder = $encoder;
    }

    protected function configure()
    {
        $this->setName('app:create-user');
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
        $user = new User();
        $user->setUsername('test');

        $plainPassword = 'test';
        $encoded = $this->encoder->encodePassword($user, $plainPassword);

        $user->setPassword($encoded);

        $this->entityManager->persist($user);
        $this->entityManager->flush();
    }
}
