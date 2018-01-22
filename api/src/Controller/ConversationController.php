<?php
declare(strict_types=1);

namespace App\Controller;

use App\Entity\Conversation;
use App\Form\Type\MessageType;
use Doctrine\ORM\EntityManagerInterface;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Method;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;
use Symfony\Component\Serializer\SerializerInterface;

class ConversationController extends Controller
{
    /**
     * @var SerializerInterface
     */
    private $serializer;

    /**
     * @var EntityManagerInterface
     */
    private $entityManager;

    /**
     * UserController constructor.
     * @param SerializerInterface $serializer
     * @param EntityManagerInterface $entityManager
     */
    public function __construct(SerializerInterface $serializer, EntityManagerInterface $entityManager)
    {
        $this->serializer = $serializer;
        $this->entityManager = $entityManager;
    }

    /**
     * @Route("/conversations/{id}")
     *
     * @param Conversation $conversation
     * @return JsonResponse
     */
    public function getConversationAction(Conversation $conversation)
    {
        return JsonResponse::fromJsonString(
            $this->serializer->serialize($conversation, 'json', ['groups' => ['conversation']])
        );
    }

    /**
     * @Route("/conversations/{id}/messages")
     * @Method({"POST"})
     *
     * @param Request $request
     * @param Conversation $conversation
     *
     * @return Response
     */
    public function createMessageAction(Request $request, Conversation $conversation)
    {
        $form = $this->createForm(MessageType::class);
        $form->submit(json_decode($request->getContent(), true));

        if ($form->isSubmitted() && $form->isValid()) {
            $message = $form->getData();
            $message->setConversation($conversation);
            $message->setFrom($this->getUser());

            $this->entityManager->persist($message);
            $this->entityManager->flush();

            return new Response('', 204);
        }

        return new Response('', 400);
    }
}
